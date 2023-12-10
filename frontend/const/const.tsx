export const FILE_UPLOAD = 'http://localhost:3000/upload/'

export interface User {
    id?: number,
    tag?: string,
    name?: string,
    surname?: string,
    amount_posts?: number,
    amount_subscribers?: number,
    amount_subscribe?: number,
    img_link?: string,
    online?: boolean
}
