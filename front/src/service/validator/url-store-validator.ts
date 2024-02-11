import { ValidationException } from "@/exceptions/validation-exception";
import { ValidatorI } from "./validator-i";

export class UrlStoreValidator implements ValidatorI {
  constructor(private url: string) {}

  /** @throws {Error} */
  validate() {
    this.validateUrl();
  }

  private validateUrl() {
    if (!this.url) {
      throw new ValidationException("必須項目です");
    }

    const pattern = /^https?:\/\/[-_.!~*\'()a-zA-Z0-9;\/?:\@&=+\$,%#]+/g;
    if (!pattern.test(this.url)) {
      throw new ValidationException("URLの形式が正しく有りません");
    }
  }
}
