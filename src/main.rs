use std::env;

/// Manages Command Line Interface for getting commands, subcommands, and flags.
mod cli;

/// Manages generating projects, build, and running
mod manager;

/// Manages the generation, serialisation, and deserialisation of config.toml
mod config;

fn main() {
    cli::process_flags(env::args());
}
