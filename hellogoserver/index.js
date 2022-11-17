/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 8000

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get('/', (req, res) => {
	res.status(200).send("Welcome to LearnCodeonline server")
})

app.get('/get', (req, res) => {
	res.status(200).json({ message: "Hello from learnCodeonline.in" })
})


app.post('/post', (req, res) => {
	// let myJson = req.body;      // your JSON
	let response = {
		"data": "i am from node server",
		"request": req.body
	}
	res.status(200).send(response);
})

app.post('/postform', (req, res) => {
	let response = {
		"data": "i am from node server post form",
		"request": req.body
	}
	res.status(200).send(JSON.stringify(response));
})


app.listen(port, () => {
	console.log(`Example app listening at http://localhost:${port}`)
})