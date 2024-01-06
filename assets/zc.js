

function submit() {
    let line = document.querySelector("#input").value
    let result = line.trim() === "" ? zcEval("drop") : zcEval(line) 
    let output = []
    for (let item of result.stack) {
        output.push(`<li>${item}</li>`)
    }
    document.querySelector("#output").innerHTML = `<ul>${output.join('\n')}</ul>`
    document.querySelector("#input").value = ""
}

window.onload = function() {
    document.querySelector("#input").onkeypress = function(evt) {
        let keyCode = evt.code || evt.key
        if (keyCode === 'Enter') {
            submit()
        }
    }
}