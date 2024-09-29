use http_body_util::{BodyExt, Full};
use hyper::{
    body::{Bytes, Incoming},
    server::conn::http1,
    service::service_fn,
    Request, Response,
};
use hyper_util::rt::TokioIo;
use once_cell::sync::Lazy;
use prometheus_client::encoding::text::encode;
use std::{io, net::SocketAddr};
use tokio::{
    net::TcpListener,
    pin,
    signal::unix::{signal, SignalKind},
};

mod registry;

/// Global registry for metrics.
pub static REGISTRY: Lazy<registry::Registry> = Lazy::new(registry::init);

/// Start a HTTP server to report metrics.
pub fn start_metrics_server(metrics_addr: SocketAddr) {
    tokio::spawn(start_metrics_server_inner(metrics_addr));
}

async fn start_metrics_server_inner(metrics_addr: SocketAddr) {
    dev_info!("Starting metrics server on {metrics_addr}");

    let tcp_listener = TcpListener::bind(metrics_addr).await.unwrap();
    let server = http1::Builder::new();
    while let Ok((stream, _)) = tcp_listener.accept().await {
        let mut shutdown_stream = signal(SignalKind::terminate()).unwrap();
        let io = TokioIo::new(stream);
        let _server = server.clone();
        tokio::task::spawn(async move {
            let conn = _server.serve_connection(
                io, service_fn(
                    move |_req: Request<Incoming>| {
                        Box::pin(async move {
                            let mut buf = String::new();
                            encode(&mut buf, &REGISTRY.registry)
                                .map_err(|e| io::Error::new(io::ErrorKind::Other, e))
                                .map(|_| {
                                    let body = Full::new(Bytes::from(buf)).boxed();
                                    Response::builder()
                                        .header(
                                            hyper::header::CONTENT_TYPE,
                                            "application/openmetrics-text; version=1.0.0; charset=utf-8",
                                        )
                                        .body(body)
                                        .unwrap()
                                })
                        })
                    })
            );
            pin!(conn);
            tokio::select! {
                _ = conn.as_mut() => {}
                _ = shutdown_stream.recv() => {
                    conn.as_mut().graceful_shutdown();
                }
            }
        });
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::net::{IpAddr, Ipv4Addr};

    #[tokio::test]
    async fn test_start_metrics_server() {
        let addr = SocketAddr::new(IpAddr::V4(Ipv4Addr::LOCALHOST), rand::random());
        start_metrics_server(addr);
        let client = reqwest::Client::new();
        let res = client.get(format!("http://{addr}")).send().await.unwrap();
        assert_eq!(res.status(), 200);
        assert!(res
            .headers()
            .get(reqwest::header::CONTENT_TYPE)
            .unwrap()
            .to_str()
            .unwrap()
            .contains("openmetrics"));
    }
}
