export {setBearerToken, getBearerToken}

const bearerTokenKeyValue = "willette_bearer_token";

function setBearerToken(bearerToken: string) {
    console.log(`calling setBearer token with value ${bearerToken}`)
    localStorage.setItem(bearerTokenKeyValue, bearerToken);
}

function getBearerToken() : string {
    const token = localStorage.getItem(bearerTokenKeyValue);
    if(token) {
        console.log(`successfully got token from localstorage: ${token}`)
        return token;
    } else {
        console.log(`failed to get token from localstorage`)
        return "";
    }
}
