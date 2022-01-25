import {getBearerToken} from "../persistence/localstorage"
import {production} from "../config"
import {logger} from "../logging";

export {getSoundcloudUrls, login, deleteSoundcloudUrl, addSoundcloudUrl, updateSoundcloudUrls}
export type {BearerToken, SoundcloudUrl, HttpResponse}

const serviceLocation = production ? "http://andrewwillette.com:9099" : "http://localhost:9099"

const getSoundcloudAllEndpoint = "/get-soundcloud-urls"
const addSoundcloudEndpoint = "/add-soundcloud-url"
const batchUpdateSoundcloudEndpoint = "/update-soundcloud-urls"
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

async function http<T>(request: RequestInfo, body: any, method: string, authorizationHeader: string): Promise<HttpResponse<T>> {
    if (body != null) {
        const opts:RequestInit = {
                method: method,
                headers: {
                    'Content-Type': 'application/json',
                    'Connection': 'keep-alive',
                    'Accept': '*/*',
                    'Accept-Encoding': 'gzip, deflate, br',
                    'Authorization': authorizationHeader
                },
                body: JSON.stringify(body)
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
        const response: HttpResponse<T> = await fetch(request).catch(reason => {
            console.log(`http fetch call failed with reason: ${reason}`)
            return Promise.reject()
        })
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
    logger.info(`Calling login with with username: ${username} , password: ${password}`)
    const data : Promise<HttpResponse<BearerToken>> = http<BearerToken>(`${serviceLocation}${loginEndpoint}`,
        {username, password}, "POST", "")
    return await data
}

/**
 * Sends DELETE request for a persisted soundcloudUrl
 * @param url
 */
async function deleteSoundcloudUrl(url: string) {
    logger.info(`Calling deleteSoundcloudUrl with url: ${url}`)
    const data: Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${deleteSoundcloudEndpoint}`,
        {url}, "DELETE", getBearerToken())
    return await data
}

async function addSoundcloudUrl(url: string) {
    logger.info(`Calling addSoundcloudUrl with url: ${url}`)
    const data : Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${addSoundcloudEndpoint}`,
        {url}, "PUT", getBearerToken())
    return await data
}

async function updateSoundcloudUrls(soundcloudUrls: SoundcloudUrl[]) {
    logger.info(`Calling updateSoundcloudUrls with soundcloudUrls: ${soundcloudUrls}`)
    const data : Promise<HttpResponse<ApiResponse>> = http<ApiResponse>(`${serviceLocation}${batchUpdateSoundcloudEndpoint}`,
        soundcloudUrls, "PUT", getBearerToken())
    return await data
}

