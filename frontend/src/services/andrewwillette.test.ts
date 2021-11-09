import {BearerToken, getSoundcloudUrls} from './andrewwillette'

test('Test getSoundcloudUrls', async () => {
    const soundcloudUrls: Promise<HttpResponse<SoundcloudUrl[]>> = getSoundcloudUrls()
    const result_1 = await soundcloudUrls;
    // console.log(result.parsedBody)
    const myBeverage = {url: 'https://soundcloud.com/user-434601011/sally-ann'};
    expect(result_1.parsedBody).toContainEqual(myBeverage);
});

test('BearerToken properties', () => {
    const token: BearerToken = {bearerToken: "testToken"}
    expect(token.bearerToken).toEqual("testToken")
});
