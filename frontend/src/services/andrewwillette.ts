export {getSoundcloudUrls};

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

export interface SoundcloudUrl {
    url: string
}

async function getSoundcloudUrls(): Promise<HttpResponse<SoundcloudUrl[]>> {
    const data : Promise<HttpResponse<SoundcloudUrl[]>> = http<SoundcloudUrl[]>(
        "http://localhost:8080/get-soundcloud-urls"
    );
    return await data;
}