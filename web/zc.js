var stackHist = []
var commandHist = []
var histPos = -1

var tabs = 0

function submit() {
    let line = document.querySelector("#input").value
    let result = ''
    histPos = -1
    if (line.trim() === "") {
        if (zcStackLen() === 0) {
            stackHist = []
            document.querySelector("#output").innerHTML = ''
            return
        }
        stackHist.push(zcStack())
        result = zcEval("drop")
    } else {
        stackHist.push(zcStack())
        result = zcEval(line)
        commandHist.unshift(line)
    }

    let output = []
    if (result.error != '') {
        result.stack = stackHist.pop() || []
        zcSetStack(result.stack)
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

function autoComplete() {
    let e = document.querySelector('#input')

    if (e.value.trim().length === 0) {
        return
    }

    let r = zcWordCompleter(e.value, e.selectionEnd)
    let common = zcCommonPrefix(r.candidates)

    var middle = ''
    if (r.candidates.length === 0) {
        tabs = 0
        return
    } else if (r.candidates.length == 1) {
        middle = r.candidates[0]
        tabs = 0
    } else {
        middle = common
        if (tabs >= 2) {
            let candidates = r.candidates.map((e) => e.replace("&", "&amp;"))
            document.querySelector("#popup").innerHTML = candidates.join(' ')
        }
    }
    e.value = r.prefix + middle + r.suffix
    e.selectionStart = e.selectionEnd = r.prefix.length + middle.length
}

function clearPopup() {
    document.querySelector("#popup").innerHTML = ''
}

function init() {
    let params = new URLSearchParams(window.location.search)
    let expr = params.get('eval')
    if (expr) {
        let e = document.querySelector("#input")
        e.value = expr
        submit()
        e.value = expr
        e.selectionStart = e.selectionEnd = expr.length
        e.focus()
    }
}

window.onload = function() {
    document.querySelector("#input").onkeypress = function(evt) {
        let keyCode = evt.code || evt.key
        if (keyCode === 'Enter') {
            submit()
        }
    }

    document.querySelector('#input').onkeydown = function(evt) {
        clearPopup()
        let keyCode = evt.code || evt.key

        if (keyCode === 'Tab') {
            tabs++
        } else {
            tabs = 0
        }

        if (keyCode === 'ArrowUp') {
            up()
            clearPopup()
        } else if (keyCode === 'ArrowDown') {
            down()
            clearPopup()
        } else if (keyCode === 'Tab') {
            autoComplete()
            evt.preventDefault()
        } else {
            clearPopup()
        }
    }

    document.querySelector('#auto').onclick = function(evt) {
        tabs++
        autoComplete()
        tabs++
        autoComplete()
        document.querySelector('#input').focus()
    }

    let e = document.querySelector('#input')
    e.focus()
}