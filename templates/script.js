document.getElementById("myForm").addEventListener("submit", function(event) {
    event.preventDefault(); // Предотвращаем стандартное поведение отправки формы

    var text = document.getElementById("text").value;
    var format = document.getElementById("format").value;

    var xhr = new XMLHttpRequest();
    xhr.open("POST", "/ascii-art", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            // document.getElementById("result").value = xhr.responseText;
            document.getElementById("pre-result").innerHTML = xhr.responseText;
        }
    };
    xhr.send("text=" + encodeURIComponent(text) + "&format=" + encodeURIComponent(format));
});