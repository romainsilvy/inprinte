
import React from 'react'

export default function BrightMode() {
    //get the background color
    const backgroundColor = document.body.style.backgroundColor;
    const color = document.body.style.color;
    alert(document.body.style.color)
    //check if the background color is #1d1c27
    if (backgroundColor === 'rgb(29, 28, 39)' || backgroundColor === '') {
        document.body.style.backgroundColor = 'white';
        document.body.style.color = 'black';
        document.getElementById('bright').style.color = 'black';
    } else {
        document.body.style.backgroundColor = '#1d1c27';
        document.body.style.color = 'white';
        document.getElementById('bright').style.color = 'white';
    }

    if (color == "") {
      document.body.style.color = 'black';
  } else {
      document.body.style.color = 'white';
  }
  return (
    <div>
    </div>
  )
}

