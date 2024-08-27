const { exec } = require("child_process");
const encoded = "aWQ="; // base64 coded 'id'
const decoded = Buffer.from(encoded, "base64").toString("utf8");

exec(decoded, (error, stdout, stderr) => {
  console.log(stdout);
});
