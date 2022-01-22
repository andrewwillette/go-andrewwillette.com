import {getBearerToken} from "../persistence/localstorage"
import {production} from "../config"

export {getSoundcloudUrls, login, deleteSoundcloudUrl, addSoundcloudUrl, updateSoundcloudUrls}
export type {BearerToken, SoundcloudUrl}

const serviceLocation = production ? "http://andrewwillette.com:9099" : "http://localhost:9099"

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const deleteSoundcloudEndpoint = "/delete-soundcloud-url"
const loginEndpoint = "/login"

/**
 * Represents response from login endpoint.
 */
interface BearerToken {
    bearerToken: string
}

interface HttpResponse<T> extends Response {
    parsedBody?: T
}

interface SoundcloudUrl {
    url: string,
    uiOrder: number
}

interface ApiResponse {
    success: boolean
}

// const requestHeaders: HeadersInit = new Headers();
// requestHeaders.set('Authorization', getBearerToken());

async function http<T>(request: RequestInfo, body: any, method: string, authorizationHeader: string): Promise<HttpResponse<T>> {
    // console.log(JSON.stringify(body))
    if (body != null) {
        const opts:RequestInit = {
                method: 'POST',
                // mode: "no-cors",
                // credentials: 'include',
                headers: {
                    'Content-Type': 'application/json',
                    'Connection': 'keep-alive',
                    'Accept': '*/*',
                    'Accept-Encoding': 'gzip, deflate, br',
                    'Authorization': authorizationHeader
                },
                body: JSON.stringify(body) // body data type must match "Content-Type" header
            }
        const response: HttpResponse<T> = await fetch(request, opts).catch(reason => {
            console.log(`http fetch call failed with reason: ${reason}`)
            return Promise.reject()
        })
        if(response.status === 201 || response.status === 200) {
            response.parsedBody = await response.json()
            .catch(exception => {
                console.log(`response.json() exception ${exception}`)
            })
        }
        return response
    } else {
        const response: HttpResponse<T> = await fetch(request)
        response.parsedBody = await response.json()
        return response
    }
}

async function getSoundcloudUrls(): Promise<HttpResponse<SoundcloudUrl[]>> {
    const data : Promise<HttpResponse<SoundcloudUrl[]>> = http<SoundcloudUrl[]>(
        `${serviceLocation}${getSoundcloudAllEndpoint}`, null, "GET", ""
    )
    return await data
}

/**
 * Sends POST login with provided credentials.
 *
 * @param username
 * @param password
 * @returns Promise<HttpResponse<BearerToken>> 
 */
async function login(username: string, password: string) {
    const data : Promise<HttpResponse<BearerToken>> = http<BearerToken>(`${serviceLocation}${loginEndpoint}`,
        {username, password}, "POST", "")
    return await data
}

/**
 * Sends DELETE request for a persisted soundcloudUrl
 * @param url
 */
async function deleteSoundcloudUrl(url: string) {
    const data: Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${deleteSoundcloudEndpoint}`,
        {url}, "DELETE", getBearerToken())
    return await data
}

async function addSoundcloudUrl(url: string) {
    const data : Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${addSoundcloudEndpoint}`,
        {url}, "PUT", getBearerToken())
    return await data
}

async function updateSoundcloudUrls(soundcloudUrls: SoundcloudUrl[]) {
    console.log("calling updateSoundcloudUrls in adrewwilet file with ")
    console.log(soundcloudUrls)
}

