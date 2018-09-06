/**
 *
 * @param val a string
 * @returns an encoded string that can be append to api url
 */
export function encode(val: string): string {
    return encodeURIComponent(val).
        replace(/%40/gi, '@').
        replace(/%3A/gi, ':').
        replace(/%24/g, '$').
        replace(/%2C/gi, ',').
        replace(/%20/g, '+').
        replace(/%5B/gi, '[').
        replace(/%5D/gi, ']');
}

/**
 * Build a URL by appending params to the end
 * @param url : the base url for the service
 * @param params : the request object. e.g. for HelloRequest would be the object of type HelloRequest
 * @returns: returns a full Url string - for GET by key/value pairs
 * @example:
 * baseUrl = "http://localhost:8080"
 * arg = {name: "wengwei", nick: "wentian"}
 * returns => http://localhost:8080?name="wengwei"&nick="wentian"
 */
export function generateQueryUrl<T>(url: string, params: T): string {
    if (!params) {
        return url;
    }

    let parts = [];


    for (let key in params) {
        let val;
        if (Object.prototype.hasOwnProperty(key)) {
            val = params[key];
        }

        if (val === null || typeof val === 'undefined') {
            return;
        }

        let k, vals;
        // if is array
        if (val.toString() === '[object Array]') {
            k = key + '[]';
        } else {
            k = key
            vals = [val];
        }

        vals.forEach(v => {
            // if is date
            if (v.toString() === '[object File]') {
                v = v.toISOString();
                // if is object
            } else if (typeof v === 'object') {
                v = JSON.stringify(v);
            }
            parts.push(encode(k) + '=' + encode(v))
        });
    }
    let serializedParams = parts.join('&');

    if (serializedParams) {
        url += (url.indexOf('?') === -1 ? '?' : '&') + serializedParams;
    }
    return url
}

/**
 *
 * @param url the base url for the service
 * @param serviceName the service name
 * @param functionName the function name
 * @example
 * baseUrl = "http://localhost:8080"
 * serviceName = "HelloService"
 * functionName = "SayHello"
 * returns => http://localhost:8080/HelloService.SayHello
 */
export function generateUrl<T>(url: string, serviceName: string, functionName: string): string {
    return url + "/" + serviceName + "."+ functionName;
}