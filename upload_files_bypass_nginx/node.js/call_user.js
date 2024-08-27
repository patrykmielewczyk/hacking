const { exec } = require("child_process");

let functions = {
  system: (command) =>
    exec(command, (error, stdout, stderr) => console.log(stdout)),
};

functions["system"]("id");
