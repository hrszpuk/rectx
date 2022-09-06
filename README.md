<h1 align="center">
  <img src="https://raw.githubusercontent.com/hrszpuk/rectx/master/rectx.png" width=250 height=250 />
</h1>
<p align="center">:radioactive: A powerful project manager for the ReCT programming language! :radioactive:</p>

<p align="center">
<a href="./LICENSE.md"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
<a href="https://github.com/hrszpuk"><img src="https://img.shields.io/github/followers/hrszpuk?style=social"></a>
<a href="https://twitter.com/hrszpuk"><img src="https://img.shields.io/twitter/follow/hrszpuk?style=social"></a>
<a href="https://github.com/hrszpuk/Owl/issues"><img src="https://img.shields.io/github/issues/hrszpuk/rectx"></a>
</p>

<p align="center">
The ReCT programming language was created by RedCubeDev, and originally a ByteSpace project.
Currently, the language has multiple compilers, packaging systems, and tools that may be intimidating for even an advanced user.
ReCTx makes your projects easier to manage by allowing you to configure everything in a easy to navigate config file!
</p>

### Features
- **Project generation**: generate project directory and settings based on different profiles to choose from.
- **Easily build and run projects**: building and running are as simple as `rectx build` or `rectx run`! 


## Installation

### Build from source
To build from source you will need cargo installed. 
Cargo will be installed a long side the Rust programming language (installation details are provided on the [official website](https://www.rust-lang.org/learn/get-started)).

To get started, first clone the project from GitHub and change directory to the project:
```
git clone https://github.com/hrszpuk/rectx.git && cd rectx
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

