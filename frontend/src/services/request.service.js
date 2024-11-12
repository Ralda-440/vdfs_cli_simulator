'use server'

export async function requestRest(type, endpoint, body) {
    const response = await fetch(`${process.env.API_URL}${endpoint}`, {
        method: type,
        headers: type === 'GET' ? undefined : {
          'Content-Type': 'application/json',
        },
        body: type === 'GET' ? undefined : JSON.stringify(body),
        cache: 'no-store',
      }
    );
    const data = await response.json();
    return data;
}