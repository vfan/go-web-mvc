import { useState, useMemo, useRef } from 'react'
import Sub from './Sub'

export default function Demo() {

  console.log("Parent Render")
  const [count, setCount] = useState(0)
  const timerRef = useRef<number>(0)

  const memoSub = useMemo(() => {
    return <Sub />
  }, [])

  function setCountByTimer() {
    if (timerRef.current) {
      clearInterval(timerRef.current)
    }
    
    
    timerRef.current = setInterval(() => {
      // console.log("Timer", count)
      setCount(currentCount => {
        console.log("Timer", currentCount) // 这里的currentCount就是最新的值
        return currentCount + 1
      })
    }, 1000)
    
    return () => {
      if (timerRef.current) {
        clearInterval(timerRef.current)
      }
    }
  }

  return (
    <div>
      <h1>Parent</h1>
      <h2>Count: {count}</h2>
      <button onClick={setCountByTimer}>Click me</button>
      <hr></hr>
      {memoSub}
    </div>
  )
}
