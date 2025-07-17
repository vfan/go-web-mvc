import React from 'react'

export   function Sub() {
console.log("Sub Render")

  return (
    <div>Sub</div>
  )
}
export default React.memo(Sub)