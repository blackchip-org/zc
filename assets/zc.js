var stackHist = []
var commandHist = []
var histPos = -1

function submit() {
    let line = document.querySelector("#input").value
    let result = ''
    if (line.trim() === "") {
        if (zcStackLen() === 0) {
            return
        }
        stackHist.push(zcStack())
        result = zcEval("drop")
    } else {
        stackHist.push(zcStack())
        result = zcEval(line)
    }
    commandHist.unshift(line)

    let output = []
    if (result.error != '') {
        result.stack = stackHist.pop() || []
    }
    let prev = []
    if (stackHist.length != 0) {
        prev = stackHist[stackHist.length - 1]
    }

    // print out previous stack
    if (prev.length > 0) {
        for (let item of prev) {
            item = annotate(item)
            output.push(`<li class='history'>${item}</li>`)
        }
        output.push(`<li class='history'>&nbsp;</li>`)
    }

    for (let i = 0; i < result.stack.length; i++) {
        let item = annotate(result.stack[i])
        let kind = (i === result.stack.length - 1) ? 'top-item' : 'stack-item'
        output.push(`<li class='${kind}'>${item}</li>`)
    }
    if (result.error !== '') {
        output.push(`<li class='error'>(!) ${result.error}</li>`)
    } else if (result.info !== '') {
        output.push(`<li class='info'>${result.info}</li>`)
    } else {
        output.push(`<li class='info'>&nbsp;</li>`)
    }

    document.querySelector("#output").innerHTML = `<ul>${output.join('\n')}</ul>`
    document.querySelector("#input").value = ""
}

function annotate(l) {
    let anno = '#!(anno)'
    let i = l.indexOf(anno)
    if ( i >= 0 ) {
        let head = l.slice(0, i)
        let tail = l.slice(i + anno.length)
        l = `${head}<span class='anno'># ${tail}</span>`
    }
    return l
}

function up() {
    if (histPos >= commandHist.length - 1) {
        return
    }
    histPos++
    document.querySelector("#input").value = commandHist[histPos]
    moveToEnd()
}

function down() {
    if (histPos <= -1) {
        return
    }
    histPos--
    var line = ''
    if (histPos !== -1) {
        line = commandHist[histPos]
    }
    document.querySelector("#input").value = line
    moveToEnd()
}

function moveToEnd() {
    let e = document.querySelector("#input")
    setTimeout(() => { e.selectionStart = e.selectionEnd = e.value.length }, 0)
}

window.onload = function() {
    document.querySelector("#input").onkeypress = function(evt) {
        let keyCode = evt.code || evt.key
        if (keyCode === 'Enter') {
            submit()
        }
    }

    document.querySelector('#input').onkeydown = function(evt) {
        let keyCode = evt.code || evt.key
        if (keyCode === 'ArrowUp') {
            up()
        } else if (keyCode === 'ArrowDown') {
            down()
        }
    }
}