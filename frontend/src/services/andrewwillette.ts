export {getSoundcloudUrls, login};
export type {BearerToken, SoundcloudUrl}

const serviceLocation = "http://localhost:8080"

async function http<T>(
    request: RequestInfo, body: any
): Promise<HttpResponse<T>> {
    if (body != null) {
        const response: HttpResponse<T> = await fetch(request, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            headers: {
                'Content-Type': 'text/plain',
                'Connection': 'keep-alive',
                'Accept': '*/*',
                'Accept-Encoding': 'gzip, deflate, br'
            },
            body: JSON.stringify(body) // body data type must match "Content-Type" header
        });
        if(response.status === 201) {
            response.parsedBody = await response.json();
        }
        return response;
    } else {
        const response: HttpResponse<T> = await fetch(request);
        response.parsedBody = await response.json();
        return response;
    }
}

interface HttpResponse<T> extends Response {
    parsedBody?: T;
}

interface SoundcloudUrl {
    url: string
}

async function getSoundcloudUrls(): Promise<HttpResponse<SoundcloudUrl[]>> {
    const data : Promise<HttpResponse<SoundcloudUrl[]>> = http<SoundcloudUrl[]>(
        `${serviceLocation}/get-soundcloud-urls`, null
    );
    return await data;
}

/**
 * Represents response from login endpoint.
 */
interface BearerToken {
    bearerToken: string
}

/**
 * Attempts login with provided credentials. Returns bearerToken if authentication is successful.
 *
 * @param username
 * @param password
 */
async function login(username: string, password: string) {
    const data : Promise<HttpResponse<BearerToken>> = http<BearerToken>(`${serviceLocation}/login`, {username, password});
    return await data;
}