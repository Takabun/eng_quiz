const app = require('express')()
const bodyParser = require('body-parser')
const { Nuxt, Builder } = require('nuxt')
const axios = require('axios')

app.use(bodyParser.json())

const apiUrl = process.env.TODO_API_URL
if (!apiUrl) {
  console.log('TODO_API_URL is undefined')
  process.exit(1)
}
console.log(`TODO_API_URL=${apiUrl}`)



app.get('/api/questions', (req, res) => {
  axios.get(`${apiUrl}/questions`)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.get('/question/api/question/:id', (req, res) => {
  const id = req.params.id
  axios.get(`${apiUrl}/question/${id}`)
    .then(resp => {
      console.log("api result GET QUESTION(ID)", resp, apiUrl)
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.get(`/question/api/answer/:id`, (req, res) => {
  const id = req.params.id
  axios.get(`${apiUrl}/answer/${id}`)
    .then(resp => {
      console.log("api result GET answer(ID)", resp, apiUrl)
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.get('/api/tags', (req, res) => {

  axios.get(`${apiUrl}/tags`)
    .then(resp => {
      console.log("api result GET TAGS", resp, apiUrl)
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

// ここからPOST
app.post('/api/tag', (req, res) => {
  // res.send(200).json(req.body)  //reqの中身を確認

  const param = {
    Name: req.body.Name
  }

  axios.post(`${apiUrl}/tag`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.post('/api/question', (req, res) => {
  const param = req.body

  axios.post(`${apiUrl}/question`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.post('/api/questionimage', (req, res) => {
  const param = req.body

  axios.post(`${apiUrl}/questionimage`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})


app.post('/api/answer', (req, res) => {
  const param = req.body

  axios.post(`${apiUrl}/answer`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})

app.post('/api/answerimage', (req, res) => {
  const param = req.body

  axios.post(`${apiUrl}/answerimage`, param)
    .then(resp => {
      res.status(resp.status).json(resp.data)
    })
    .catch(err => {
      res.send({ err })
    })
})




let config = require('./nuxt.config.ts')
config.dev = !(process.env.NODE_ENV === 'production')
const nuxt = new Nuxt(config)

if (config.dev) {
  const builder = new Builder(nuxt)
  builder.build()
}

app.use(nuxt.render)
app.listen(3000)
