import type { SentenceCollectionItem } from "@nuxt/content";

export class SentenceDocImpl {
  readonly sentence: string;
  readonly description: string;
  readonly gherkinExample: string;
  readonly variables: SentenceVariable[];
  constructor(_data: SentenceCollectionItem) {
    const data = unref(_data);

    this.variables = [];
    console.log(data.variables);

    // for (const element of data.value.variables) {
    //   console.log(element);

    //   //   this.variables.push({
    //   //     name: element.name,
    //   //     type: element.type,
    //   //     description: element.description,
    //   //     example: element.example,
    //   //   });
    // }

    this.sentence = data.sentence;
    this.description = data.description;

    // this.variables = [];
    this.gherkinExample = "";
  }
}

export type SentenceDoc = {
  readonly sentence: string;
  readonly description: string;
  readonly gherkinExample: string;
  readonly variables: SentenceVariable[];
};

export type SentenceVariable = {
  name: string;
  type: string;
  description?: string;
  example?: string;
};
