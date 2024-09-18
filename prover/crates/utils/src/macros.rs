/// This macro is used to notify sp1 cycle tracker that a new routine has started.
#[macro_export]
macro_rules! cycle_track {
    ($e:expr, $($arg:tt)*) => {
        {
            #[cfg(all(feature = "sp1", feature = "cycle-tracker"))]
            println!("cycle-tracker-start: {}", format!($($arg)*));

            #[allow(clippy::let_and_return)]
            let __cycle_track_result = $e;

            #[cfg(all(feature = "sp1", feature = "cycle-tracker"))]
            println!("cycle-tracker-end: {}", format!($($arg)*));

            __cycle_track_result
        }
    };
}

/// This macro is used to notify sp1 cycle tracker that a new routine has started.
#[macro_export]
macro_rules! cycle_tracker_start {
    ($($arg:tt)*) => {
        #[cfg(all(feature = "sp1", feature = "cycle-tracker"))]
        println!("cycle-tracker-start: {}", format!($($arg)*));
    };
}

/// This macro is used to notify sp1 cycle tracker that a routine has ended.
#[macro_export]
macro_rules! cycle_tracker_end {
    ($($arg:tt)*) => {
        #[cfg(all(feature = "sp1", feature = "cycle-tracker"))]
        println!("cycle-tracker-end: {}", format!($($arg)*));
    };
}

/// This macro is for logging level trace
#[macro_export]
macro_rules! dev_trace {
    ($($arg:tt)*) => {
        #[cfg(any(feature = "dev", test))]
        {
            $crate::tracing::trace!($($arg)*);
        }
    };
}

/// This macro is for logging level info
#[macro_export]
macro_rules! dev_info {
    ($($arg:tt)*) => {
        #[cfg(any(feature = "dev", test))]
        {
            $crate::tracing::info!($($arg)*);
        }
    };
}

/// This macro is for logging level error
#[macro_export]
macro_rules! dev_error {
    ($($arg:tt)*) => {
        #[cfg(any(feature = "dev", test))]
        {
            $crate::tracing::error!($($arg)*);
        }
    };
}

/// This macro is for logging level debug
#[macro_export]
macro_rules! dev_debug {
    ($($arg:tt)*) => {
        #[cfg(any(feature = "dev", test))]
        {
            $crate::tracing::debug!($($arg)*);
        }
    };
}

/// This macro is for logging level warn
#[macro_export]
macro_rules! dev_warn {
    ($($arg:tt)*) => {
        #[cfg(any(feature = "dev", test))]
        {
            $crate::tracing::warn!($($arg)*);
        }
    };
}

/// This macro is for measuring duration to metrics
#[macro_export]
macro_rules! measure_duration_histogram {
    ($label:ident, $e:expr) => {{
        #[cfg(feature = "metrics")]
        let __measure_duration_histogram_start = std::time::Instant::now();

        #[allow(clippy::let_and_return)]
        let __measure_duration_histogram_result = $e;

        #[cfg(feature = "metrics")]
        $crate::metrics::REGISTRY
            .$label
            .observe(__measure_duration_histogram_start.elapsed().as_millis() as f64);

        #[cfg(feature = "metrics")]
        dev_debug!(
            "measured duration {} = {:?}",
            stringify!($label),
            __measure_duration_histogram_start.elapsed(),
        );

        __measure_duration_histogram_result
    }};
}

/// This macro is for update gauge to metrics
#[macro_export]
macro_rules! update_metrics_gauge {
    ($label:ident, $e:expr) => {
        #[cfg(feature = "metrics")]
        {
            $crate::metrics::REGISTRY.$label.set($e);
        }
    };
}

/// This macro is for update counter to metrics
#[macro_export]
macro_rules! update_metrics_counter {
    ($label:ident) => {
        #[cfg(feature = "metrics")]
        {
            $crate::metrics::REGISTRY.$label.inc();
        }
    };
}
