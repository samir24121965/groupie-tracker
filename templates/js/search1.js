let searchBox = document.getElementById('search-input');

searchBox.addEventListener('keyup', search);


let timeoutId;

function search(e) {

    clearTimeout(timeoutId);

    let artistName = document.getElementById('search-artist');
    let artistList = document.getElementsByClassName('opt-artist');
    if (artistList) {
        Array.from(artistList).forEach(elm => elm.remove())
    }
    let creationDateStart = document.getElementById('creationdatestart').value
    // alert(creationDateStart)
    timeoutId = setTimeout(() => {
        let keywords = e.target.value.toLowerCase().trim();
        index = keywords.indexOf("-")
        if (index > -1) {
            keywords = keywords.substring(0, index).trim();
        }
        let elements = document.getElementsByClassName('main-article');
        let authorMemb;

        Array.from(elements).forEach((elm) => {
            // name = elm.getAttribute('data-name');
            allNames = elm.getAttribute('data-allNames').toLowerCase();
            authorMemb = elm.getAttribute('data-authorMemb');
            creationDate = parseInt(elm.getAttribute('data-creationDate'));
            // alert("hello" + creationDate + typeof creationDate);

            if (allNames.includes(keywords)) {
                let names = allNames.split("|").filter(elm => elm.includes(keywords));
                // the name of band is the same of the author
                names = [...new Set(names)] // eliminate duplicates in an array
                for (name of names) {
                    elm.style.display = "block";
                    let opt = document.createElement('option');
                    opt.value = capitalize(name + " - " + authorMemb);
                    opt.className = "opt-artist"
                    artistName.appendChild(opt);
                    
                }
            } else {
                elm.style.display = "none";
            }
        });
    }, 500);
}

function capitalize(str) {
    return str.replace(/(^\w{1})|(\s+\w{1})/g, letter => letter.toUpperCase());
}