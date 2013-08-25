function distance(left, right) {
  var result = Math.max(left.length, right.length)
  var rightMap = mapLetters(right)
  for (var x = 0; x < left.length; x++) {
    var locations = rightMap[left[x]]
    if (typeof(locations) != "undefined") {
      for (var i = 0; i < locations.length; i++) {
        var y = locations[i]
        var beforeMatch = Math.max(x, locations[i])
        if (beforeMatch < result) {
          var xOut = x
          var yOut = y
          while (xOut < left.length && yOut < right.length &&
                 left[xOut] == right[yOut]) {
            xOut++
            yOut++
          }
          var afterMatch = distance(left.substring(xOut), right.substring(yOut))
          result = Math.min(result, beforeMatch + afterMatch)
        }
      }
    }
  }
  return result
}
      
function mapLetters(text) {
  var result = {}
  result.keys = ""
  for (var i = 0; i < text.length; i++) {
    var letter = text[i]
    if (result.keys.indexOf(letter) == -1) {
      result[letter] = []
      result.keys += letter
    }
    result[letter].push(i)
  }
  return result
}

// command-line interface

if (typeof window === 'undefined') {
  var fs = require('fs')
  
  if (process.argv.length < 4) {
    console.log("USAGE: node levenshtein.js /path/to/dict numberOfWords")
    process.exit(1)
  }

  var text = fs.readFileSync(process.argv[2], { encoding : "UTF-8" })
  var words = []

  var i = text.indexOf("\n")
  while (i != -1) {
    words.push(text.substr(0, i))
    text = text.substr(i + 1)
    i = text.indexOf("\n")
  }

  var limit = parseInt(process.argv[3])
  var count = 0
  console.time("crunching")
  for (var i = 0; i < limit - 1; i++) {
    for (var j = i + 1; j < limit; j++) {
      console.log(words[i] + " vs. " + words[j] + ": " + distance(words[i], words[j]))
    }
  }
  console.timeEnd("crunching")
  console.log(" for " + (limit * (limit - 1) / 2) + " pairs")
}
