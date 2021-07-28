export {getAllSoundcloudUrls};
export type {SoundcloudUrls};

async function http<T>(
    request: RequestInfo
): Promise<HttpResponse<T>> {
    const response: HttpResponse<T> = await fetch(request);
    response.parsedBody = await response.json();
    return response;
}

interface HttpResponse<T> extends Response {
    parsedBody?: T;
}
// represents json with list of soundcloud urls
interface SoundcloudUrls {
    Urls: string[];
}

async function getAllSoundcloudUrls(): Promise<HttpResponse<SoundcloudUrls>> {
    const data : Promise<HttpResponse<SoundcloudUrls>> = http<SoundcloudUrls>(
        "http://localhost:8080/get-soundcloud-urls"
    );
    return await data;
}