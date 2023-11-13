export const addHoursToDate = (d: Date, h: number) => {
  d.setTime(d.getTime() + h * 60 * 60 * 1000);
  return d;
};
