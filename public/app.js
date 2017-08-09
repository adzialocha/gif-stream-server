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

const container = document.getElementById('container')

function update() {
  request('/api/stream')
    .then(response => {
      response.data.forEach(animation => {
        const img = document.createElement('img')
        img.src = animation.url

        container.appendChild(img)
      })
    })
}

update()
