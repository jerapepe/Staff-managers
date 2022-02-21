package main

import (
    "log"
    "net/http"
    "html/template"
)

func main() {
    runServer()
    //readExcel()
    // a := getVal("Hoja2", "A1")
    // b := getVal("Hoja2", "A3")
    // c := getVal("Hoja2", "G3")
    // fmt.Println("Dia:")
    // d := getVal("Hoja2", "A4")
    // fmt.Println(a)
}

type Empleado struct {
    dicx map[int]string
}
type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos []Todo
    Dpps []string
}

func hello(w http.ResponseWriter, req *http.Request) {  
        data := getRowss()

        ida := searchDat(data, "ID:")
        idsa := addDat(ida, 4)
	    idss := searName(idsa, data)

        id := searchDat(data, "Nombre:")
        ids := addDat(id, 2)
	    names := searName(ids, data)
        
        //empleado := Empleado{}
        //arregloempleado := []Empleado{}
        var dicx map[int]string
        dicx = map[int]string{}
        // var dic map[int]string
        dic := make(map[int]interface{})
        Dats := make(map[string]interface{})
        
        
        for x:=0; x<len(idss); x++{
            dic[x] = idss[x]
            dicx[x] = names[x]
            Dats["id"] = dic[x]
            //log.Println("key=",x,"value=",names[x])
        }
        
        
        //Dats["name"] = dicx 
        
        user := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 1", Done: false},
            },
            Dpps: Dats,
        }

        tmpl := template.Must(template.ParseFiles("layout.html"))
        tmpl.Execute(w, user)
    }

func runServer() {
    http.HandleFunc("/hello", hello)
    log.Println("Server run...")
    http.ListenAndServe(":8090", nil)
}

// [2 3 4 5 6 8 9 10 11 12 13 15 16 17 18 31 32 33 35 36 37 39 41 42 43 44 45 46 47 49 50 52 53 59 60 63 64 65 67 68 69 71 73 74 75 76 77 78 79 81 82 85 86 87 89 90 91 93 95 96 97 98 99 100 101 103 104 106 107 112 113 114 117 118 119 121 122 123 125 127 128 129 130 131 132 133 135 136 138 139 145 146 149 150 151 153 154 155 157 159 160 161 162 163 164 165 167 168 171 172 173 175 176 177 179 181 182 183 184 185 186 187 189 190 192 193 199 200 203 204 205 207 208 209 211 213 214 215 216 217 218 219 221 222 224 225 228 231 232 235 236 237 239 240 241 243 245 246 247 248 249 250 251 253 254 256 
// 263 264 267 268 269 271 272 273 275 277 278 279 280 281 282 283 285 286 288 289 295 296 299 300 301 303 304 305 307 309 310 311 312 313 314 315 317 318 320 321 323 326 327 328 330 331 332 334 336 337 338 339 340 341 342 344 345 347 348 354 355 358 359 360 362 363 364 366 368 369 370 371 372 373 374 376 377 379 380 386 387 390 391 392 394 395 396 398 400 401 402 403 404 405 406 408 409 411 412 418 419 422 423 424 426 427 428 430 432 433 434 435 436 437 438 440 441 443 451 452 453 455 456 
// 457 459 461 462 463 464 465 466 467 469 470 472 473 474 475 476 477 480 481 482 484 485 486 488 490 491 492 493 494 495 496 498 499 501 502 508 509 512 513 514 516 517 518 520 522 523 524 525 526 527 528 530 531 533 534 536 540 541 544 545 546 548 549 550 552 554 555 556 557 558 559 560 562 563 565 566 572 573 576 577 578 580 581 582 584 586 587 588 589 590 591 592 594 595 597 598 601 604 605 608 609 610 612 613 614 616 618 619 620 621 622 623 624 626 627 629 630 636 637 640 641 642 644 
// 645 646 648 650 651 652 653 654 655 656 658 659 661 662 668 669]
