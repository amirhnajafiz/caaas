// url
const url = 'http://' + window.location.host;

// make http call for register
function register() {
    let stEl = document.getElementById("student-number");
    let psEl = document.getElementById("password");

    let data = {
        "student_number": stEl.value,
        "password": psEl.value
    }

    fetch(url+"/api/register", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    })
        .then((response) => response.text())
        .then((data) => {
            document.getElementById("response").innerText = data
        })
        .catch((error) => {
            console.error(error);
            alert("Registration Failed");
        })
}

function copy() {
    let text = document.getElementById("response").innerText

    navigator.clipboard.writeText(text)
        .catch((error) => {
            console.log(error)
        });
}