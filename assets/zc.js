var hist = []

function submit() {
    let line = document.querySelector("#input").value
    let result = '' 
    if (line.trim() === "") {
        if (zcStackLen() === 0) {
            return 
        }
        hist.push(zcStack())
        result = zcEval("drop") 
    } else {
        hist.push(zcStack())
        result = zcEval(line) 
    }

    let output = []
    if (result.error != '') {
        result.stack = hist.pop() || []
    }
    let prev = []
    if (hist.length != 0) {
        prev = hist[hist.length - 1]
    }

    // print out previous stack 
    if (prev.length > 0) {
        for (let item of prev) {
            output.push(`<li class='history'>${item}</li>`)
        }
        output.push(`<li class='history'>&nbsp;</li>`)
    }

    for (let i = 0; i < result.stack.length; i++) {
        let item = result.stack[i]
        let kind = (i === result.stack.length - 1) ? 'top-item' : 'stack-item'
        output.push(`<li class='${kind}'>${item}</li>`)
    }
    if (result.error !== '') {
        output.push(`<li class='error'>(!) ${result.error}</li>`)
    } else if (result.info !== '') {
        output.push(`<li class='info'>${result.info}</li>`)
    } else {
        output.push('<li>&nbsp;</li>')
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