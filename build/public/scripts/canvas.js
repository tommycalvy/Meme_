var container = document.querySelector('#stage-parent');
var width = container.offsetWidth;
var height = container.offsetHeight;

var stage = new Konva.Stage({
  container:    'container',
  width:         width,
  height:        height
});
var layer = new Konva.Layer();
stage.add(layer);
let searchByImageUrl = "https://www.google.com/searchbyimage?hl=en-US&image_url=";
let baseUrl = "https://meme-242116.appspot.com";


// Finding the source image url from drag start event
document.addEventListener('dragstart', function(e) {
  //console.log(e.target)
  //console.log(e.target.src)
  $.ajax(baseUrl + '/image-search?url=' + e.target.src)
  .done(data => console.log(data))
  .fail((xhr, status) => console.log('error:', status));


  e.dataTransfer.effectAllowed = 'copy';
  console.log(e);
  console.log(e.target);
  console.log(e.target.src);
  e.dataTransfer.setData('text/plain', e.target.src);
});

var con = stage.container();
con.addEventListener('dragover', function(e) {
  e.preventDefault(); // !important
});

con.addEventListener('drop', function(e) {
  e.preventDefault();
  if (e.stopPropagation) {
    e.stopPropagation(); // Stops some browsers from redirecting.
  }
  // now we need to find pointer position
  // we can't use stage.getPointerPosition() here, because that event
  // is not registered by Konva.Stage
  // we can register it manually:
  stage.setPointersPositions(e);

  Konva.Image.fromURL(e.dataTransfer.getData('text/plain'), function(image) {
    layer.add(image);

    image.position(stage.getPointerPosition());
    image.draggable(true);

    layer.draw();
  });
});

window.addEventListener('resize', function() {
  width = container.offsetWidth;
  height = container.offsetHeight;
  stage.width(width);
  stage.height(height);
  stage.draw();
});

stage.on('click tap', function(e) {
  // if click on empty area - remove all transformers
  if (e.target === stage) {
    stage.find('Transformer').destroy();
    layer.draw();
    return;
  }

  // remove old transformers
  // TODO: we can skip it if current rect is already selected
  stage.find('Transformer').destroy();

  // create new transformer
  var tr = new Konva.Transformer();
  layer.add(tr);
  tr.attachTo(e.target);
  layer.draw();
});

stage.on('transformstart', function() {
  console.log('transform start');
});
