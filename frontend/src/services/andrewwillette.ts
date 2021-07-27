export {getAllSoundcloudUrls}

async function http<T>(
    request: RequestInfo
): Promise<T> {
    const response = await fetch(request);
    return await response.json();
}

interface SoundcloudUrls {
    Urls: string[];
}

function getAllSoundcloudUrls() {
    let soundcloudUrls: Array<string>;

    const data = http<SoundcloudUrls[]>(
        "http://localhost:8080/get-soundcloud-urls"
    );

    data.then((result) => {
        console.log(result);
        return result;
    });
}