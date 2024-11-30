export function randomAvatar() {
  const randomNumber = Math.floor(Math.random() * 90000) + 1
  return `https://avatars.githubusercontent.com/u/${randomNumber}?v=4`
}
