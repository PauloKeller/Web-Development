const users = [
	{
		id: "1",
		name: "Paulo Keller",
		email: "test@test.com",
		age: 28
	},
	{
		id: "2",
		name: "vinicius Keller",
		email: "test@test.com",
		age: 28
	},
	{
		id: "3",
		name: "Keller",
		email: "test@test.com",
		age: 28
	}
]

const posts = [
	{
		id: "10",
		title: "post1",
		body: "body1",
		published: true,
		author: "1"
	},
	{
		id: "11",
		title: "post4",
		body: "body1",
		published: false,
		author: "1"
	},
	{
		id: "12",
		title: "post3",
		body: "body2",
		published: true,
		author: "2"
	}
]

const comments = [
	{
		id: "100",
		text: "awoidjioawjd",
		author: '1',
		post: "10"
	},
	{
		id: "101",
		text: "awopdkawjd",
		author: '2',
		post: "11"
	},
	{
		id: "102",
		text: "awkdpoawkdopk",
		author: '3',
		post: "12"
	},
]

const db = {
  users,
  posts,
  comments
}

export { db as default } 