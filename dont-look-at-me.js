const fs = require('fs')
const https = require('https')

// Fetches emoji data
function getData() {
  return new Promise(resolve => {
    https.get('https://raw.githubusercontent.com/muan/emojilib/master/emojis.json', (res) => {
      let data = ''
      res.on('data', chunk => data += chunk)
      res.on('end', () => {
        console.log(data)
        resolve(JSON.parse(data))
      })
    })
  })
}

// make array JSON with set keys
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
