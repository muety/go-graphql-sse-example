// https://stackoverflow.com/a/1214753/3112139
function addMinutes (date, minutes) {
  return new Date(date.getTime() + minutes * 60000)
}

// https://stackoverflow.com/a/7709819/3112139
// TODO: i18n
function formatDiff (date1, date2, nonNegative, short, locale) {
  const diffMs = Math.max((date1 - date2), nonNegative ? 0 : Number.MIN_SAFE_INTEGER)
  const diffDays = Math.floor(diffMs / 86400000)
  const diffHrs = Math.floor((diffMs % 86400000) / 3600000)
  const diffMins = Math.round(((diffMs % 86400000) % 3600000) / 60000)

  const displayMins = x => short ? 'm' : (x === 1 ? 'Minute' : 'Minuten')
  const displayHrs = x => short ? 'h' : (x === 1 ? 'Stunde' : 'Stunden')
  const displayDays = x => short ? 'd' : (x === 1 ? 'Tag' : 'Tage')

  let str = ''
  if (diffDays) str += `${diffDays} ${displayDays(diffDays)}, `
  if (diffHrs) str += `${diffHrs} ${displayHrs(diffHrs)}, `
  str += `${diffMins} ${displayMins(diffMins)}`

  return str
}

export {
  addMinutes,
  formatDiff
}
