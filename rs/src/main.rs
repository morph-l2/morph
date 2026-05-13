use clap::{App, AppSettings, Arg, SubCommand};

use vsock_sample::command_parser::{ClientArgs, ServerArgs};
use vsock_sample::create_app;
use vsock_sample::{client, server};

fn main() {
    let app = create_app!();
    let args = app.get_matches();

    match args.subcommand() {
        Some(("server", args)) => {
            let server_args = ServerArgs::new_with(args).unwrap();
            server(server_args).unwrap();
        }
        Some(("client", args)) => {
            let client_args = ClientArgs::new_with(args).unwrap();
            client(client_args).unwrap();
        }
        Some(_) | None => ()
    }
}
