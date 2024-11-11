'use server'

export async function requestRest(type, endpoint, body) {
    //console.log(process.env.API_URL,"este es el api url");
    const response = await fetch(`${process.env.API_URL}/${endpoint}`, {
        method: type,
        headers: {
          'Content-Type': 'application/json',
        },
        body: type === 'GET' ? undefined : JSON.stringify(body),
      }
    );
    return response.json();
}