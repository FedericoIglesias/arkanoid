const canvas = document.querySelector('canvas')
const ctx = canvas.getContext('2d')

const $sprite = document.querySelector('#sprite')
const $bricks = document.querySelector('#bricks')

canvas.width = 448
canvas.height = 400

// game variable 

// Object in the canvas
//BALL
const ballRadius = 4
let x = canvas.width / 2
let y = canvas.height - 30
let dx = 2
let dy = -2

// Paddle
const paddleHeight = 10
const paddleWidth = 50

let paddlex = (canvas.width - paddleWidth) / 2
let paddley = canvas.height - paddleHeight - 10

let rightPressed = false
let leftPressed = false

// Bricks

const brickRowCount = 6
const brickColumnCount = 13
const brickWidth = 32
const brickHeight = 15
const brickPadding = 0
const brickOffSetTop = 80
const brickOffsetLeft = 16
const bricks = []

const BRICK_STATUS = {
  ACTIVE: 1,
  DESTROY: 0
}

for (let c = 0; c < brickColumnCount; c++) {
  bricks[c] = []
  for (let r = 0; r < brickRowCount; r++) {
    const brickX = c * (brickWidth + brickPadding) + brickOffsetLeft
    const brickY = r * (brickHeight + brickPadding) + brickOffSetTop
    const random = Math.floor(Math.random() * 8)
    bricks[c][r] = {
      x: brickX,
      y: brickY,
      status: BRICK_STATUS.ACTIVE,
      color: random
    }
  }
}

// DRAW

function drawBall() {
  ctx.beginPath()
  ctx.arc(x, y, ballRadius, 0, Math.PI * 2)
  ctx.fillStyle = '#fff'
  ctx.fill()
  ctx.closePath()
}

function drawPaddle() {
  ctx.drawImage(
    $sprite, // img
    29, //where star cut the img: clip x
    174, //where star cut the img: clip y
    paddleWidth, //size the cut X
    paddleHeight, //size the cut Y
    paddlex, // position X
    paddley,  // position Y
    paddleWidth, // width draw
    paddleHeight // height draw
  )
}


function drawBricks() {
  for (let c = 0; c < brickColumnCount; c++) {
    for (let r = 0; r < brickRowCount; r++) {
      const currentBrick = bricks[c][r]
      if (currentBrick.status === BRICK_STATUS.DESTROY)
        continue;

      const clipX = currentBrick.color * 32

      ctx.drawImage(
        $bricks,
        clipX,
        0,
        32,
        14,
        currentBrick.x,
        currentBrick.y,
        brickWidth,
        brickHeight
      )

    }
  }
}


// Moves
function collisionDetection() {
  for (let c = 0; c < brickColumnCount; c++) {
    for (let r = 0; r < brickRowCount; r++) {
      const currentBrick = bricks[c][r]
      if (currentBrick.status === BRICK_STATUS.DESTROY) {
        continue
      }

      const isBallSameXAsBrick =
        x > currentBrick.x &&
        x < currentBrick.x + brickWidth

      const isBallSameYAsBrick =
        y + dy > currentBrick.y &&
        y < currentBrick.y + brickHeight

      if (isBallSameXAsBrick && isBallSameYAsBrick) {
        dy = -dy
        currentBrick.status = BRICK_STATUS.DESTROY  
      }
    }
  }
}

function ballMovement() {
  if (x + dx > canvas.width - ballRadius || //Right Wall
    x + dx < 0 + ballRadius)// Left Wall
  { dx = -dx }

  if (y + dy < ballRadius // top
  ) { dy = -dy }
  x += dx
  y += dy

  const isBallSameXAsPaddle =
    x > paddlex && x < paddlex + paddleWidth

  const isBallSameYAsPaddle = y + dy > paddley

  if (isBallSameXAsPaddle && isBallSameYAsPaddle) {
    dy = -dy
  } else if (y + dy > canvas.height) {
    document.location.reload()
  }
}
function paddleMovement() {
  if (rightPressed && paddlex < canvas.width - paddleWidth) {
    paddlex += 7
  } else if (leftPressed && paddlex > 0) {
    paddlex -= 7
  }
}

function cleanCanvas() {
  ctx.clearRect(0, 0, canvas.width, canvas.height)
}

function initEvents() {
  document.addEventListener('keydown', keyDownHandler)
  document.addEventListener('keyup', keyUpHandler)

  function keyDownHandler(event) {
    const { key } = event
    if (key === 'Right' || key === 'ArrowRight') {
      rightPressed = true
    } else if (key === 'Left' || key === 'ArrowLeft') {
      leftPressed = true
    }
  }
  function keyUpHandler(event) {
    const { key } = event
    if (key === 'Right' || key === 'ArrowRight') {
      rightPressed = false
    } else if (key === 'Left' || key === 'ArrowLeft') {
      leftPressed = false
    }
  }
}

function draw() {
  cleanCanvas()
  drawBall()
  drawPaddle()
  drawBricks()

  collisionDetection()
  ballMovement()
  paddleMovement()


  window.requestAnimationFrame(draw)
}

draw()
initEvents()
