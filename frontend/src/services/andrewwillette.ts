import {getBearerToken} from "../persistence/localstorage"
import {production} from "../config"

export {getSoundcloudUrls, login, deleteSoundcloudUrl, addSoundcloudUrl}
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
    url: string
}


interface ApiResponse {
    success: boolean
}

async function http<T>(
    request: RequestInfo, body: any
): Promise<HttpResponse<T>> {
    if (body != null) {
        console.log("body is not null")
        const response: HttpResponse<T> = await fetch(request, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            headers: {
                'Content-Type': 'text/plain',
                'Connection': 'keep-alive',
                'Accept': '*/*',
                'Accept-Encoding': 'gzip, deflate, br'
            },
            body: JSON.stringify(body) // body data type must match "Content-Type" header
        }).catch(reason => {
            console.log(`post fetch reason is ${reason}`)
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
        `${serviceLocation}${getSoundcloudAllEndpoint}`, null
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
    const data : Promise<HttpResponse<BearerToken>> = http<BearerToken>(`${serviceLocation}${loginEndpoint}`, {username, password})
    return await data
}

async function deleteSoundcloudUrl(url: string) {
    const data: Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${deleteSoundcloudEndpoint}`, {url, bearerToken: getBearerToken()})
    return await data
}

async function addSoundcloudUrl(url: string) {
    const data : Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${addSoundcloudEndpoint}`, {url, bearerToken: getBearerToken()})
    return await data
}

