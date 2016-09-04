(function() {

  var COMMON_WORD_COUNT = 3000;

  window.app = {};

  function init() {
    window.app.fetchDictionary('assets/tv.csv', function(err, dict) {
      if (err) {
        var loader = document.getElementById('loader');
        loader.textContent = 'Failed to load: ' + err.toString();
      } else {
        gotDictionary(dict);
      }
    });
  }

  function gotDictionary(dict) {
    document.body.className = '';
    var input = document.getElementById('text-input');

    var renderTimeout = null;
    input.addEventListener('keydown', function() {
      if (renderTimeout) {
        return;
      }
      renderTimeout = setTimeout(function() {
        renderTimeout = null;
        renderFrequencies(dict);
      }, 100);
    });
  }

  function renderFrequencies(dict) {
    var input = document.getElementById('text-input');
    var freqView = document.getElementById('freq-view');
    freqView.innerHTML = '';

    var words = input.value.split(/\s/);
    for (var i = 0, len = words.length; i < len; ++i) {
      var word = words[i];
      var ranking = (COMMON_WORD_COUNT*2 - dict.ranking(word)) /
        (COMMON_WORD_COUNT * 2);
      if (ranking < 0) {
        ranking = 0;
      }
      var color;
      if (ranking > 0.5) {
        color = 'rgba(123, 186, 209, ' + ((ranking-0.5)*2).toFixed(2) + ')';
      } else {
        color = 'rgba(212, 113, 143, ' + (1 - ranking*2).toFixed(2) + ')';
      }
      console.log(color);
      var wordView = document.createElement('span');
      wordView.className = 'word-view';
      wordView.textContent = word;
      wordView.style.backgroundColor = color;
      freqView.appendChild(wordView);
    }
  }

  window.addEventListener('load', init);

})();
