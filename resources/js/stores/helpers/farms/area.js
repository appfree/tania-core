export const AreaTypes = [
  { key: 'SEEDING',  label: 'Seeding' },
  { key: 'GROWING', label: 'Growing Area' }
]

export const AreaLocations = [
  { key: 'OUTDOOR',  label: 'Field (Outdoor)' },
  { key: 'INDOOR', label: 'Greenhouse (Indoor)' }
]

export const AreaSizeUnits = [
  { key: 'Ha',  label: 'Ha' },
  { key: 'm2', label: 'meter square' }
]

export function FindAreaType(key) {
  return AreaTypes.find(item => item.key === key)
}

export function FindAreaLocation(key) {
  return AreaLocations.find(item => item.key === key)
}

export function FindAreaSizeUnit(key) {
  return AreaSizeUnits.find(item => item.key === key)
}
