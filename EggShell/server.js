var express = require('express');
var app = express();
var router = express.Router();


var rtIndex = require('./routes/Index.js');

app.use(router)


//routes
app.use('/', rtIndex);


app.listen(8080);

module.exports.router = router;
