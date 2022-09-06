# rectx
A quick and easy to use project manager for ReCT programming language based projects.

The ReCT programming language was created by RedCubeDev, and originally a ByteSpace project.
Currently, the language is multiple compilers, package systems, and projects that may be intimidating.
ReCTx tries to simplify projects by automatically setting up everything you will need in your ReCT journey!

### Features
- **Project generation**: generate project directory and settings based on different profiles to choose from.
- **Easily build and run projects**: building and running are as simple as `rectx build` or `rectx run`! 


## Installation

### Build from source
To build from source you will need cargo installed. 
Cargo will be installed a long side the Rust programming language (installation details are provided on the [official website](https://www.rust-lang.org/learn/get-started)).

To get started, first clone the project from GitHub and change directory to the project:
```
git clone https://github.com/hrszpuk/rectx.git
cd rectx
```
From here, use cargo to build the project in release mode:
```
cargo build --release
```
The executable (named `rectx`) will be generated within `./target/release`.
You can move the executable into an installation path such as `/usr/local/bin` or your own directory which is on the `PATH`.
```
sudo mv ./target/release/rectx /usr/local/bin
```


## Usage

### Creating new projects
To create a new project using `rectx`, use the `new` command as shown below.
```
rectx new "project-name"
```
This will generate a folder with the name `"project-name"`.
The project directory will contain a `README.md` with a title of the same name as the directory, a `.gitignore` file, and a `/src` directory which houses a `main.rct`.
The `main.rct` file will contain a simple hello world program.

### Building a project
Ensure you are within the project directory and run the command shown below.
```
rectx build
```
This will build the source file `main.rct` in `/src` into an exectuable called `main`.
You can find `main` in `/src` and run it using `./src/main`

### Running a project
While you're in the project directory you can run the `rectx run` command shown below.
```
rectx run
```
This will build the soruce file `main.rct` in `/src` into an exectuable called `main`.
This command will also run the executable, you should be able to see any output in the console.

