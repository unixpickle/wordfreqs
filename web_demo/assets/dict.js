(function() {

  var READY_STATE_DONE = 4;
  var HTTP_OK = 200;

  function Dict(dictCSV) {
    this._dictionary = {};
    this._maxFreq = 0;
    var lines = dictCSV.split('\n');
    for (var i = 0, len = lines.length; i < len; ++i) {
      if (!lines[i]) {
        continue;
      }
      var entry = lines[i].split(',');
      if (entry.length !== 3) {
        throw new Error('unexpected number of CSV columns: ' + entry.length);
      }
      var entryInfo = {
        rank: parseInt(entry[1]),
        freq: parseFloat(entry[2])
      };
      this._dictionary[entry[0].toLowerCase()] = entryInfo;
      this._maxFreq = Math.max(this._maxFreq, entryInfo.freq);
    }
  }

  Dict.prototype.ranking = function(word) {
    var trimmed = word.replace(/(^["'\(]|["'\)\.;,\?]$)/, '');
    trimmed = trimmed.toLowerCase();
    if (!this._dictionary.hasOwnProperty(trimmed)) {
      return Infinity;
    }
    var entry = this._dictionary[trimmed];
    return entry.rank;
  };

  function fetchDictionary(url, cb) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url);
    xhr.send(null);
    xhr.onreadystatechange = function() {
      if (xhr.readyState === READY_STATE_DONE) {
        if (xhr.status === HTTP_OK) {
          try {
            cb(null, new Dict(xhr.responseText));
          } catch (e) {
            cb(e, null);
          }
        } else {
          cb('HTTP status: '+xhr.status, null);
        }
      }
    };
  }

  window.app.fetchDictionary = fetchDictionary;

})();
