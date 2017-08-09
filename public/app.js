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

window.setInterval(() => {
  request('http://localhost:3000/api/stream')
}, 10000)
