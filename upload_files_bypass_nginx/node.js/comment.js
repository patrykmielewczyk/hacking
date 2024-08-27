const { exec } = require("child_process");

// comment: /*id*/
let command = "id";
exec(command, (error, stdout, stderr) => {
  console.log(stdout);
});
