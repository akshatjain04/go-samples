import fs from "fs";
import csvParser from "csv-parser";

function convertToNumber(value) {
  const intValue = parseInt(value);
  if (!isNaN(intValue) && value.trim() !== "") {
    return intValue;
  }

  const floatValue = parseFloat(value);
  if (!isNaN(floatValue) && value.trim() !== "") {
    return floatValue;
  }

  return value;
}

function readCsvData(filePath) {
  return new Promise((resolve, reject) => {
    const results = [];
    fs.createReadStream(filePath)
      .pipe(csvParser())
      .on("data", (data) => {
        const convertedData = {};
        Object.keys(data).forEach((key) => {
          convertedData[key] = convertToNumber(data[key]);
        });
        results.push(convertedData);
      })
      .on("end", () => resolve(results))
      .on("error", (err) => reject(err));
  });
}

export { readCsvData }
