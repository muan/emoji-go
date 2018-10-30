const fs = require('fs')
const https = require('https')

// Fetches emoji data
function getData() {
  return new Promise(resolve => {
    https.get('https://raw.githubusercontent.com/muan/emojilib/master/emojis.json', (res) => {
      let data = ''
      res.on('data', chunk => data += chunk)
      res.on('end', () => {
        resolve(JSON.parse(data))
      })
    })
  })
}

// Makes array JSON with set keys
// Because I've concluded Go hates arbitrary JSON hates Go hates me
async function makeABetterJSON() {
  const data = await getData()

  const betterData = []
  Object.keys(data).forEach(function(key) {
    const newSet = data[key]
    newSet["name"] = key
    betterData.push(newSet)
  })

  fs.writeFileSync("./emoji-for-go.json", JSON.stringify(betterData))
}

makeABetterJSON()
