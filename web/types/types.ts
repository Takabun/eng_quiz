interface Tag {
  id: number,
  name: string
}

interface Image {
  url: string;
}

interface Question {
  id: number,
  created_at: Date,
  user: string,
  text: string,
  default_image: number,
  tags: Tag[],
  images: Image[]
}

interface Raw_Tag {
  ID: number,
  Name: string
}

interface Raw_Image {
  Url: string;
}

interface Raw_Question {
  ID: number,
  CreatedAt: Date,
  UpdatedAt: Date,
  User: string,
  Text: string,
  DefaultImage: number,
  Tags: Raw_Tag[],
  QuestionImages: Raw_Image[]
}


export {Tag, Image, Question, Raw_Tag, Raw_Image, Raw_Question}