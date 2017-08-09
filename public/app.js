function request(url, data, method = 'GET') {
  return new Promise((resolve, reject) => {
    const xhr = new XMLHttpRequest()
    xhr.open(method, url)
    xhr.onreadystatechange = () => {
      if (xhr.readyState === 4) {
        if (xhr.status === 200) {
          resolve(JSON.parse(xhr.responseText))
        } else {
          reject()
        }
      }
    }
    xhr.send(data)
  })
}

const streamElem = document.getElementById('stream')
const loadButtonElem = document.getElementById('load-button')

let nextMarker

function update() {
  const path = nextMarker ? `/api/stream?marker=${nextMarker}` : '/api/stream'

  loadButtonElem.disabled = true

  request(path)
    .then(response => {
      if (response.nextMarker) {
        loadButtonElem.disabled = false
        nextMarker = response.nextMarker
      }

      response.data.forEach(animation => {
        const img = document.createElement('img')
        img.src = animation.url

        streamElem.appendChild(img)
      })
    })
}

loadButtonElem.addEventListener('click', update)
