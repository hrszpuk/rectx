/**  Cli module
*   This module manages the command line interface of rectx!
*   Features:
*   - Handling program arguments (commands and options)
*   - Displaying help menus
*   - Recognising commands and calling the correct functions
 **/
use std::{fs, io};
use std::io::{ErrorKind, stdout, Write};
use std::process::exit;
use colored::Colorize;
use crate::manager;
use crate::manager::generate_project_directory;

/// Takes the command line arguments and calls the correct function.
pub fn process_flags(args: std::env::Args) {

    // Moving the program arguments into a custom vector.
    // We do this so we can remove some parts of the program arguments
    // such as the application name.
    let mut arguments = Vec::new();
    for arg in args {
        arguments.push(arg);
    }

    // Drop application name ./rectx
    arguments.remove(0);

    // Here we just match for different commands
    // and call the correct function.
    // Example: "new" -> new_project()
    if arguments.len() > 0 {
        match arguments[0].as_str() {
            "help" => help(),
            "new" => new_project(&arguments),
            "build" => build_project(false),
            "run" => build_project(true),
            _ => help_unknown()
        }
    } else {
        // Here we call the help menu if we only encounter "rectx"
        help()
    }
}

/// Creates a new project.
/// This is called when the user calls the command "new".
pub fn new_project(args: &Vec<String>) {
    if args.len() > 1 {

        process(format!("Creating a new project named \"{}\"", args[1]));

        // Checking for errors
        match generate_project_directory(&args[1]) {

            Ok(()) => success(format!("Created new project \"{}\"", args[1])),
            Err(error) => match error.kind() {

                // Will occur if "main.rs" or "readme" fail to be created and
                // the write_all method is still triggered. This is unlikely,
                // but is handled anyway (just in case ofc).
                ErrorKind::NotFound => abort(
                    String::from("A previously generated file could not be found!")
                ),

                // Will occur if rectx doesn't have permission to generate project files or folders
                // where the user requests it to be generated.
                // A common cause of this could be user attempting to create a new project in
                // administrator only directories (such as / on *nix).
                ErrorKind::PermissionDenied => abort(
                    String::from(
                        "Attempted to generate project entities but received permission denied!"
                    )
                ),

                // Will occur if a directory of this name already exists.
                // Here we ask if the user would like to overwrite the directory,
                // and if not we simply abort
                ErrorKind::AlreadyExists => {
                    issue(format!(
                        "The directory \"{}\" already exists and therefore could not be generated!",
                        args[1],
                    ));

                    // Here we ask about overwriting the directory
                    let overwrite = question(format!("Would you like to overwrite the directory \"{}\"?", args[1]));
                    if overwrite {

                        // Of course overwriting could also error, in that case we abort
                        process(String::from("Attempting to remove directory!"));
                        match fs::remove_dir_all(&args[1]) {

                            // Overwriting successful, we retry to create project
                            Ok(()) => {
                                success(format!("Removed directory \"{}\"!", args[1]));
                                process(String::from("Retrying project creation..."));
                                generate_project_directory(&args[1]).expect("!!! Project creation failed again... aborting!");
                                success(format!("Created new project \"{}\"", args[1]))
                            },

                            // Overwriting unsuccessful, quickly abort
                            Err(_) => abort(
                                String::from("Attempt to overwrite directory failed!")
                            )
                        }
                    } else {

                        // This is if the user doesn't want to overwrite btw
                        abort(
                            String::from("Failed to generate project because directory already exists!")
                        );
                    }
                },
                // Will occur if no more memory is avaiable on the user's machine
                // as we allocate memory in generate_project_directory
                ErrorKind::OutOfMemory => abort(
                    format!(
                        "Could not generate project files for \"{}\" because machine is out of memory!",
                        args[1],
                    )
                ),
                _ => abort(
                    String::from(
                        "Unable to recover from an unexpected error while generating new project!",
                    )
                )
            }
        };
    } else {
        help_new();
    }
    exit(0);
}

/// Builds the project in the current directory.
/// This is called when the user calls the command "build".
/// This also handles running the executable with argument "run"
pub fn build_project(run: bool) {

    process(String::from("Building project executable"));

    // Checking for errors
    match manager::generate_project_executable(run) {

        Ok(()) => success(String::from("Executable was successful!")),
        Err(error) => match error.kind() {
            ErrorKind::NotFound => abort(
                String::from("Could not generate executable as source file(s) could not be found!")
            ),
            ErrorKind::PermissionDenied => abort(
                String::from("Could not generate executable due to lack of permissions!")
            ),
            ErrorKind::OutOfMemory => abort(
                String::from("Could not generate executable as machine is out of memory!")
            ),
            _ => abort(
                String::from("Unexpected unrecoverable error occurred while generating project executable!")
            ),
        },
    }
    exit(0);
}

/// Use cli::success when a successful process has taken place
pub fn success(message: String) {
    println!(
        "[SUCCESS] {}", message
    );
}

/// Use cli::process when a process has began (logging)
pub fn process(message: String) {
    println!(
        "::: {}", message
    );
}

/// use cli::issue when an issue is found, but can be recovered
pub fn issue(message: String) {
    println!(
        ":!: {}", message
    );
}

/// Use cli::abort when an unrecoverable issue is found (exits the program)
pub fn abort(message: String) {
    println!(
        "!!! {}", message
    );
    println!(
        "[ABORT] An unrecoverable error caused rectx to abort!"
    );
    panic!();
}

/// Use cli::info when wanting to give the user non-specific information
#[allow(dead_code)]
pub fn info(message: String) {
    println!(
        "[INFO] {}", message
    );
}

/// Use cli::question when wanting to get information from the user
pub fn question(message: String) -> bool {
    let mut buffer = String::new();

    print!("{} [Y/n] ", message);
    stdout().flush().unwrap();

    match io::stdin().read_line(&mut buffer) {
        Ok(_) => (),
        Err(_) => abort(
            format!(
                "Attempted to question user but found unrecoverable error instead!"
            )
        )
    }

    match buffer.to_lowercase().trim() {
        "y"|"yeah"|"yes"|"ye"|"yeet"|"yee" => true,
        _ => false
    }
}

/// Show help menu to the user.
/// The help menu contains information on commands and flags, and what they do.
pub fn help() {
    let title = "ReCTx Project Manager";
    let version = "v1.0.0";
    let mut dashes = String::new();
    for _ in 0..(title.len()+version.len()+2) {
        dashes.push('-');
    }
    println!("{}\n{} {}\n{}\n{} {}\n\n{}",
             dashes, title.bold(), version.green(),
             dashes, "Usage:".bold(),
             "rectx <command>".bright_green(),
             "[Commands]".bold());
    for (name, example, description) in help_commands() {
        println!(
            "{}\t {} {}\t {} {}",
            name.bold().bright_cyan(),
            ":".black(),
            example.bold().bright_blue(),
            ":".black(),
            description.bold().bright_green(),
        )
    }

    println!("\n{}{}",
             "For more information check out the GitHub page: ".bold(),
             "https://github.com/hrszpuk/rectx".bright_blue());
    println!("{} {}",
    "or join the ReCT Discord server:".bold(),
    "https://discord.gg/Ymm9xGxWZf".bright_blue());
    exit(0);
}

pub fn help_commands() -> Vec<(&'static str, &'static str, &'static str)> {
    vec![
        ("Help", "rectx help", "Shows this help message."),
        ("New", "rectx new", "Creates a new ReCTx project."),
        ("Build", "rectx build", "Builds the current ReCTx project."),
        ("Run", "rectx run", "Runs the current ReCTx project.")
    ]
}

/// A specific help menu for the "new" command.
/// Creating a new project using the new command: usage and explanation.
pub fn help_new() {
    println!(
        "ReCTx :: Help Menu :: \"new\"

Usage: rectx new project-name

This command will create a new ReCT project with the name provided.
The project will contain: /src/main.rs, README.md, and config.toml!

For more information visit the GitHub page: https://github.com/hrszpuk/rectx"
    )
}

pub fn help_unknown() {
    println!(
        "ReCTX :: Help Menu :: Unknown

Usage: rectx <command> [options]

The command you have entered does not seem to exist!
Use \"rectx help\" for more information on the commands you can use.

For more information visit the GitHub page: https://github.com/hrszpuk/rectx"
    )
}