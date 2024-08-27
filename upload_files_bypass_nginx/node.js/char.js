const { exec } = require("child_process");

let cmd = String.fromCharCode(105) + String.fromCharCode(100); // ASCII for 'id'
exec(cmd, (error, stdout, stderr) => {
  console.log(stdout);
});
