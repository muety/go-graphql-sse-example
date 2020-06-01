import { formatDiff } from '../util/date'

function filter (value, reverse, short) {
  if (!(value instanceof Date)) {
    value = new Date(value)
  }

  if (reverse) return formatDiff(new Date(), value, true, short)
  return formatDiff(value, new Date(), true, short)
}

export default filter
