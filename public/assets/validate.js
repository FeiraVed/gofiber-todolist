document.querySelector(".formAdd").addEventListener("submit",(event)=>{
  const input = document.querySelector("input")
  const re = /^[a-zA-Z ]+$/
  console.info(input.value.trim().match(re))
  if (input.value.trim().match(re) == null) {
    const newElement = document.createElement("span")
    newElement.textContent = "Permission Denied"
    newElement.style.color = "red"
    newElement.style.display = "block"
    newElement.style.fontFamily = "Anonymous Pro"
    document.querySelector(".message").appendChild(newElement)
    event.preventDefault()
  }
})
