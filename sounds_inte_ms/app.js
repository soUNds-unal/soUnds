const express = require('express');
const app = express();
const router = express.Router();

const path = __dirname + '/src/views/';
const port = 3302;

router.use(function (req,res,next) {
  console.log('/' + req.method);
  next();
});

router.get('/', function(req,res){
  res.sendFile(path + 'index.html');
});

router.get('/sharks', function(req,res){
  res.sendFile(path + 'sharks.html');
});

router.get('/playlist', function(req,res){
  res.sendFile(path + 'playlist.html');
});

app.use(express.static(path));
app.use('/', router);

app.listen(port, function () {
  console.log('Example app listening on port 3302!')
})
