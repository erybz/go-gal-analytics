var BASE_API_URL = "https://go-gal.herokuapp.com/stats"

$(document).ready(function () {
    loadStats("country", "Country", "#country-table", "#country-chart")
    loadStats("city", "City", "#city-table", "#city-chart")
    loadStats("deviceType", "Device", "#device-type-table", "#device-type-chart")
    loadStats("devicePlatform", "Platform", "#device-platform-table", "#device-platform-chart")
    loadStats("os", "OS", "#device-os-table", "#device-os-chart")
    loadStats("browser", "Browser", "#device-browser-table", "#device-browser-chart")
    loadStats("language", "Language", "#device-language-table", "#device-language-chart")
    loadStats("referral", "Referral", "#referral-table", "#referral-chart")
});

function loadStats(statsName, title, tableID, chartID) {
    $.ajax({
        url: BASE_API_URL + "?q=" + statsName,
        success: function (response) {
            $('#total-page-views').text(totalPageViews(response))
            createNewTable(statsName, response, title, tableID)
            createNewChart(statsName, response, chartID)
        }
    });
}


function createNewTable(statsName, response, tableTitle, tableID) {
    response.sort(function (a, b) {
        return b.pageViews - a.pageViews
    })

    var tableHead = "<thead><tr><th>" + tableTitle + "</th><th>Page Views</th></tr></thead>"
    $(tableID).append(tableHead)
    var tableBody = "<tbody>"
    for (var k in response) {
        if (response[k][statsName] == "") {
            response[k][statsName] = "Unknown"
        }
        tableBody += "<tr><td>" + response[k][statsName] + "</td><td>" + response[k]["pageViews"] + "</td></tr>"
    }
    tableBody += "</tbody>"
    $(tableID).append(tableBody)
}

function createNewChart(query, queryResponse, elemID) {
    var datasetLabel = "Page Views"

    var label = []
    var data = []

    var countryLabels = []
    var countryData = []
    var cityLabels = []
    var cityData = []
    var deviceTypeLabels = []
    var deviceTypeData = []
    var devicePlatformLabels = []
    var devicePlatformData = []
    var osLabels = []
    var osData = []
    var browserLabels = []
    var browserData = []
    var languageLabels = []
    var languageData = []
    var referralLabels = []
    var referralData = []

    var backgroundColor = [
        'rgba(54, 162, 235, 0.2)',
        'rgba(54, 162, 235, 0.2)',
        'rgba(54, 162, 235, 0.2)',
        'rgba(54, 162, 235, 0.2)',
        'rgba(54, 162, 235, 0.2)',
    ]
    var borderColor = [
        'rgba(54, 162, 235, 1)',
        'rgba(54, 162, 235, 1)',
        'rgba(54, 162, 235, 1)',
        'rgba(54, 162, 235, 1)',
        'rgba(54, 162, 235, 1)',
    ]

    var chartOptions = {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true
                }
            }]
        }
    };

    queryResponse.sort(function (a, b) {
        return b.pageViews - a.pageViews
    })

    if (queryResponse.length > 10) queryResponse.length = 10;

    for (var k in queryResponse) {
        switch (query) {
            case "country":
                if (queryResponse[k].country == "") {
                    queryResponse[k].country = "Unknown"
                }
                countryLabels.push(queryResponse[k].country)
                countryData.push(queryResponse[k].pageViews)
                label = countryLabels
                data = countryData
                break;
            case "city":
                if (queryResponse[k].city == "") {
                    queryResponse[k].city = "Unknown"
                }
                cityLabels.push(queryResponse[k].city)
                cityData.push(queryResponse[k].pageViews)
                label = cityLabels
                data = cityData
                break;
            case "deviceType":
                if (queryResponse[k].deviceType == "") {
                    queryResponse[k].deviceType = "Unknown"
                }
                deviceTypeLabels.push(queryResponse[k].deviceType)
                deviceTypeData.push(queryResponse[k].pageViews)
                label = deviceTypeLabels
                data = deviceTypeData
                break;
            case "devicePlatform":
                if (queryResponse[k].devicePlatform == "") {
                    queryResponse[k].devicePlatform = "Unknown"
                }
                devicePlatformLabels.push(queryResponse[k].devicePlatform)
                devicePlatformData.push(queryResponse[k].pageViews)
                label = devicePlatformLabels
                data = devicePlatformData
                break;
            case "os":
                if (queryResponse[k].os == "") {
                    queryResponse[k].os = "Unknown"
                }
                osLabels.push(queryResponse[k].os)
                osData.push(queryResponse[k].pageViews)
                label = osLabels
                data = osData
                break;
            case "browser":
                if (queryResponse[k].browser == "") {
                    queryResponse[k].browser = "Unknown"
                }
                browserLabels.push(queryResponse[k].browser)
                browserData.push(queryResponse[k].pageViews)
                label = browserLabels
                data = browserData
                break;
            case "language":
                if (queryResponse[k].language == "") {
                    queryResponse[k].language = "Unknown"
                }
                languageLabels.push(queryResponse[k].language)
                languageData.push(queryResponse[k].pageViews)
                label = languageLabels
                data = languageData
                break;
            case "referral":
                if (queryResponse[k].referral == "") {
                    queryResponse[k].referral = "Unknown"
                }
                referralLabels.push(queryResponse[k].referral)
                referralData.push(queryResponse[k].pageViews)
                label = referralLabels
                data = referralData
                break;
        }
    }

    var myChart = new Chart($(elemID), {
        type: 'bar',
        data: {
            labels: label,
            datasets: [{
                label: datasetLabel,
                data: data,
                barPercentage: 0.5,
                barThickness: 20,
                maxBarThickness: 30,
                minBarLength: 2,
                borderWidth: 1,
                backgroundColor: backgroundColor,
                borderColor: borderColor
            }]
        },
        options: chartOptions
    });

}

function totalPageViews(response) {
    var views = 0
    for (k in response) {
        views += response[k].pageViews
    }
    return views
}